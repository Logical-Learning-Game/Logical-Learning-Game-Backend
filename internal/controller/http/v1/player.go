package v1

import (
	"fmt"
	"llg_backend/internal/dto"
	"llg_backend/internal/entity/nullable"
	"llg_backend/internal/pkg/httputil"
	"llg_backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	mapConfigService service.MapConfigurationService
	statisticService service.PlayerStatisticService
}

func NewPlayerController(mapConfigService service.MapConfigurationService, statisticService service.PlayerStatisticService) *PlayerController {
	return &PlayerController{
		mapConfigService: mapConfigService,
		statisticService: statisticService,
	}
}

func (c PlayerController) initRoutes(handler *gin.RouterGroup) {
	h := handler.Group("/players")
	{
		playerGroup := h.Group("/:playerID")
		{
			playerGroup.GET("/available_maps", c.ListAvailableMaps)
			playerGroup.POST("/statistics", c.CreateSessionHistory)
		}
	}
}

func (c PlayerController) ListAvailableMaps(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	playerWorlds, err := c.mapConfigService.ListFromPlayerID(ctx, playerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	worldDTOs := make([]*dto.WorldDTO, 0, len(playerWorlds))
	for _, world := range playerWorlds {
		mapConfigDTOs := make([]*dto.MapConfigurationDTO, 0, len(world.MapConfigurationForPlayers))
		for _, mapConfig := range world.MapConfigurationForPlayers {
			pqInt32Tile := mapConfig.MapConfiguration.Tile

			intTile := make([]int, len(pqInt32Tile))
			for i := range pqInt32Tile {
				intTile[i] = int(pqInt32Tile[i])
			}

			mapConfiguration := mapConfig.MapConfiguration

			rules := make([]*dto.RuleDTO, 0, len(mapConfiguration.Rules))
			for _, rule := range mapConfiguration.Rules {
				pqInt32Parameter := rule.Parameters
				intParameter := make([]int, len(pqInt32Parameter))
				for i := range pqInt32Parameter {
					intParameter[i] = int(pqInt32Parameter[i])
				}

				ruleDTO := &dto.RuleDTO{
					MapRuleID:  rule.ID,
					RuleName:   rule.RuleName,
					Theme:      rule.Theme,
					Parameters: intParameter,
				}

				rules = append(rules, ruleDTO)
			}

			mapConfigDTO := &dto.MapConfigurationDTO{
				MapID:                      mapConfiguration.ID,
				MapName:                    mapConfiguration.ConfigName,
				Tile:                       intTile,
				Height:                     int(mapConfig.MapConfiguration.Height),
				Width:                      int(mapConfig.MapConfiguration.Width),
				MapImagePath:               nullable.NullString{NullString: mapConfiguration.MapImagePath},
				Difficulty:                 mapConfiguration.Difficulty,
				StarRequirement:            int(mapConfiguration.StarRequirement),
				LeastSolvableCommandGold:   int(mapConfiguration.LeastSolvableCommandGold),
				LeastSolvableCommandSilver: int(mapConfiguration.LeastSolvableCommandSilver),
				LeastSolvableCommandBronze: int(mapConfiguration.LeastSolvableCommandBronze),
				LeastSolvableActionGold:    int(mapConfiguration.LeastSolvableActionGold),
				LeastSolvableActionSilver:  int(mapConfiguration.LeastSolvableActionSilver),
				LeastSolvableActionBronze:  int(mapConfiguration.LeastSolvableActionBronze),
				Rules:                      rules,
				IsPass:                     mapConfig.IsPass,
				TopHistory:                 nil,
			}

			mapConfigDTOs = append(mapConfigDTOs, mapConfigDTO)
		}

		worldDTO := &dto.WorldDTO{
			WorldID:   world.ID,
			WorldName: world.Name,
			Maps:      mapConfigDTOs,
		}

		worldDTOs = append(worldDTOs, worldDTO)
	}

	ctx.JSON(http.StatusOK, worldDTOs)
}

func (c PlayerController) CreateSessionHistory(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	var createSessionHistoryRequestDTO dto.CreateGameSessionRequestDTO
	if err := ctx.ShouldBindJSON(&createSessionHistoryRequestDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	gameSession, err := c.statisticService.CreateSessionHistory(ctx, playerID, createSessionHistoryRequestDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.Header("Location", fmt.Sprintf("/v1/sessions/%d", gameSession.ID))
	ctx.Status(http.StatusCreated)
}
