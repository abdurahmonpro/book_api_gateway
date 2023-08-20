package handlers

import (
	"api_gateway/api/http"
	"api_gateway/genproto/auth_service"
	"api_gateway/genproto/book_service"
	"api_gateway/models"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
// @ID create_book
// @Router /books [POST]
// @Summary Create Book
// @Description  Create Book
// @Tags Book
// @Accept json
// @Produce json
// @Param profile body book_service.CreateBook true "CreateBookRequestBody"
// @Success 200 {object} http.Response{data=book_service.BookResponse} "GetBookBody"
// @Response 400 {object} auth_service.UserWrongResponse "book not found"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateBook(c *gin.Context) {
	var Book book_service.CreateBook

	err := c.ShouldBindJSON(&Book)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.BookService().Create(
		c.Request.Context(),
		&Book,
	)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, &auth_service.UserWrongResponse{
			Data:    "openlibrary did not responded propperly",
			IsOk:    true,
			Message: "book not found",
		})
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetBookByTitle godoc
// @ID get_book_by_title
// @Router /books/:title [GET]
// @Summary Get Book By title
// @Description Get Book By title
// @Tags Book
// @Accept json
// @Produce json
// @Param title path string true "title"
// @Success 200 {object} http.Response{data=book_service.BookResponse} "BookBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetBookByTitle(c *gin.Context) {
	title := c.Param("title")

	resp, err := h.services.BookService().GetBookByTitle(
		context.Background(),
		&book_service.BookByTitle{
			Title: title,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetBookList godoc
// @ID get_book_list
// @Router /books [GET]
// @Summary Get Book List
// @Description Get Book List
// @Tags Book
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=book_service.BookResponse} "BookResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetBookList(c *gin.Context) {

	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.BookService().GetList(
		context.Background(),
		&book_service.BookListRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// PatchCBook godoc
// @ID patch_book
// @Router /books/{id} [PATCH]
// @Summary Patch Book
// @Description Patch Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body models.UpdatePatchRequest true "UpdatePatchRequestBody"
// @Success 200 {object} http.Response{data=book_service.BookResponse} "Book data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdatePatchBook(c *gin.Context) {
	var updatePatchBook models.UpdatePatchRequest

	err := c.ShouldBindJSON(&updatePatchBook)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	bookid := c.Param("id")

	intbookid, err := strconv.Atoi(bookid)

	resp, err := h.services.BookService().UpdatePatch(
		c.Request.Context(),
		&book_service.UpdatePatchBook{
			Id: int32(intbookid),
			// Status: updatePatchBook.Status, // Use the new 'Status' field
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteBook godoc
// @ID delete_book
// @Router /books/{id} [DELETE]
// @Summary Delete Book
// @Description Delete Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=book_service.BookResponse} "Book data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteBook(c *gin.Context) {
	BookId := c.Query("id")

	intbookid, err := strconv.Atoi(BookId)
	if err != nil {
		fmt.Println("****************5***********")
		fmt.Println("****************5***********")
		fmt.Println("****************5***********")
		fmt.Println("****************5***********")
		return
	}

	resp, err := h.services.BookService().Delete(
		c.Request.Context(),
		&book_service.BookPK{Id: int32(intbookid)},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
