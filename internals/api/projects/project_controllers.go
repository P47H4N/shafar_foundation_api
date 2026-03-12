package projects

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/P47H4N/shafar_foundation_api/internals/helpers"
	"github.com/gin-gonic/gin"
)

func InitProjectControllers(ps *ProjectServices) *ProjectControllers {
	return &ProjectControllers{
		service: ps,
	}
}

func (pc *ProjectControllers) GetProjects(c *gin.Context) {
	projects, err := pc.service.GetProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: fmt.Sprintf("%d projects found.", len(projects)),
		Data: projects,
	})
}

func (pc *ProjectControllers) GetProjectsById(c *gin.Context) {
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid project id.",
		})
		return
	}
	project, err := pc.service.GetProjectsById(uint(paramId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: "Project found.",
		Data: project,
	})
}

func (pc *ProjectControllers) CreateProjects(c *gin.Context) {
	var projectBody CreateProjectBody
	if err := c.ShouldBindBodyWithJSON(&projectBody); err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid Data",
		})
		return
	}
	if err := pc.service.CreateProjects(&projectBody); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: "Project Created Successfully.",
	})
}

func (pc *ProjectControllers) UpdateProject(c *gin.Context) {
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid project id.",
		})
		return
	}
	var updateProject CreateProjectBody
	if err := c.ShouldBindBodyWithJSON(&updateProject); err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid Data",
		})
		return
	}
	if err := pc.service.UpdateProject(&updateProject, uint(paramId)); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: "Project Updated Successfully",
		Data: updateProject,
	})
}

func (pc *ProjectControllers) DeleteProject(c *gin.Context) {
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid project id.",
		})
		return
	}
	if err := pc.service.DeleteProject(uint(paramId)); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: "Project Deleted Successfully.",
	})
}
