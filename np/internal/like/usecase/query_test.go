package usecase_test

import (
	"context"
	"newsportal/internal/like/dto"
	"newsportal/internal/like/entity"
	"newsportal/internal/like/mocks"
	"newsportal/internal/like/usecase"

	eAuth "newsportal/internal/authorization/entity"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type deps struct {
	SessionProvider   *mocks.MockSessionProvider
	QueryRepository   *mocks.MockQueryRepository
	CommandRepository *mocks.MockCommandRepository
}

type body struct {
	requestAPI    dto.CountRequest
	requestToRepo []usecase.CountRequest
	response      []usecase.CountResponse
}

func sessionGenerator() {}

func TestQueryUseCase(t *testing.T) {
	tableTest := []struct {
		name      string
		setup     func(d *deps)
		request   dto.CountRequest
		response  []usecase.CountResponse
		assertErr func(t *testing.T, err error)
	}{
		{
			name: "error_UserID_equal_to_zero",
			setup: func(d *deps) {
			},
			request:  dto.CountRequest{},
			response: nil,
			assertErr: func(t *testing.T, err error) {
				targetErr := usecase.UserIDEqualToZero
				require.ErrorIs(t, err, targetErr)
			},
		},

		{
			name: "error_request_target_len_zero",
			setup: func(d *deps) {
			},
			request: dto.CountRequest{
				UserID: 1,
			},
			response: nil,
			assertErr: func(t *testing.T, err error) {
				targetErr := usecase.TargetLengthIsEqualToZero
				require.ErrorIs(t, err, targetErr)
			},
		},

		{
			name: "error_session_user_does_not_exist",
			setup: func(d *deps) {
				d.SessionProvider.
					EXPECT().
					Session(gomock.Any(), 999999999999).
					Return(nil, eAuth.ErrUserDoesNotExist)
			},
			request: dto.CountRequest{
				UserID: 999999999999,
				Targets: []dto.Target{
					{TargetID: 1, TargetType: "article"},
					{TargetID: 2, TargetType: "article"},
				},
			},
			response: nil,
			assertErr: func(t *testing.T, err error) {
				targetErr := eAuth.ErrUserDoesNotExist
				require.ErrorIs(t, err, targetErr)
			},
		},

		{
			name: "error_by_some_reason_dont_have_a_permission",
			setup: func(d *deps) {
				d.SessionProvider.
					EXPECT().
					Session(gomock.Any(), 21001).
					Return(&eAuth.Session{
						UserID: 21001,
						Roles:  nil,
						MergedPermission: map[eAuth.Domain]map[eAuth.Permission]bool{
							eAuth.Like: {
								eAuth.CanRead: false,
							},
						},
					}, nil)
			},
			request: dto.CountRequest{
				UserID: 21001,
				Targets: []dto.Target{
					{TargetID: 1, TargetType: "article"},
					{TargetID: 2, TargetType: "article"},
				},
			},
			response: nil,
			assertErr: func(t *testing.T, err error) {
				var targetErr *usecase.PermissionError
				require.ErrorAs(t, err, &targetErr)
			},
		},

		{
			name: "error_non_valid_targetID",
			setup: func(d *deps) {
				d.SessionProvider.
					EXPECT().
					Session(gomock.Any(), 21001).
					Return(&eAuth.Session{
						UserID: 21001,
						Roles:  nil,
						MergedPermission: map[eAuth.Domain]map[eAuth.Permission]bool{
							eAuth.Like: {
								eAuth.CanRead: true,
							},
						},
					}, nil)
			},
			request: dto.CountRequest{
				UserID: 21001,
				Targets: []dto.Target{
					{TargetID: -999, TargetType: "article"},
				},
			},
			response: nil,
			assertErr: func(t *testing.T, err error) {
				var targetErr *usecase.ValidationError
				require.ErrorAs(t, err, &targetErr)
			},
		},

		{
			name: "error_non_valid_targetType",
			setup: func(d *deps) {
				d.SessionProvider.
					EXPECT().
					Session(gomock.Any(), 21001).
					Return(&eAuth.Session{
						UserID: 21001,
						Roles:  nil,
						MergedPermission: map[eAuth.Domain]map[eAuth.Permission]bool{
							eAuth.Like: {
								eAuth.CanRead: true,
							},
						},
					}, nil)
			},
			request: dto.CountRequest{
				UserID: 21001,
				Targets: []dto.Target{
					{TargetID: 10000, TargetType: "nonValidTargetType"},
				},
			},
			response: nil,
			assertErr: func(t *testing.T, err error) {
				var targetErr *usecase.ValidationError
				require.ErrorAs(t, err, &targetErr)
			},
		},

		{
			name: "happy_path",
			setup: func(d *deps) {
				d.SessionProvider.
					EXPECT().
					Session(gomock.Any(), 21001).
					Return(&eAuth.Session{
						UserID: 21001,
						Roles:  nil,
						MergedPermission: map[eAuth.Domain]map[eAuth.Permission]bool{
							eAuth.Like: {
								eAuth.CanRead: true,
							},
						},
					}, nil)

				request := []usecase.CountRequest{
					{usecase.Target{
						ID:   10000,
						Type: "article",
					}},
				}

				repoResponse := []usecase.CountResponse{
					{
						Target: usecase.Target{
							ID:   10000,
							Type: "article",
						},
						Count: map[entity.ReactionType]int{
							entity.Positive: 10,
							entity.Negative: 15,
						},
					},
				}

				d.QueryRepository.
					EXPECT().
					CountReaction(gomock.Any(), request).
					Return(repoResponse, nil)

			},

			request: dto.CountRequest{
				UserID: 21001,
				Targets: []dto.Target{
					{TargetID: 10000, TargetType: "article"},
				},
			},

			response: []usecase.CountResponse{
				{
					Target: usecase.Target{
						ID:   10000,
						Type: "article",
					},
					Count: map[entity.ReactionType]int{
						entity.Positive: 10,
						entity.Negative: 15,
					},
				},
			},

			assertErr: func(t *testing.T, err error) {
				var targetErr error = nil
				require.ErrorIs(t, err, targetErr)
			},
		},
	}

	for _, concreteTest := range tableTest {
		t.Run(concreteTest.name, func(t *testing.T) {
			// arrange
			var (
				ctx  context.Context    = t.Context()
				ctrl *gomock.Controller = gomock.NewController(t)
				d    *deps              = &deps{
					SessionProvider:   mocks.NewMockSessionProvider(ctrl),
					QueryRepository:   mocks.NewMockQueryRepository(ctrl),
					CommandRepository: mocks.NewMockCommandRepository(ctrl),
				}
				err error
			)

			if concreteTest.setup != nil {
				concreteTest.setup(d)
			}

			uc := usecase.NewQueryUseCase(&usecase.Deps{
				SessionProvider:   d.SessionProvider,
				QueryRepo:         d.QueryRepository,
				CommandRepository: d.CommandRepository,
			})

			// act
			actual, err := uc.CountReaction(ctx, concreteTest.request)

			// assert
			concreteTest.assertErr(t, err)
			require.ElementsMatch(t, concreteTest.response, actual)
		})
	}

}
