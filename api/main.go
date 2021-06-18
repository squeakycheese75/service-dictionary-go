package main

import (
	"fmt"

	"net/http"

	router "github.com/squeakycheese75/service-dictionary-go/api/http"

	"github.com/squeakycheese75/service-dictionary-go/api/controller"
	"github.com/squeakycheese75/service-dictionary-go/api/repository"
	"github.com/squeakycheese75/service-dictionary-go/api/service"
)

var (
	sourceRepository repository.SourceRepository = repository.NewSqlLiteSourceRepository("test.db")
	sourceService    service.SourceService       = service.NewSourceService(sourceRepository)
	sourceController controller.SourceController = controller.NewSourceController(sourceService)
	httpRouter       router.Router               = router.NewMuxRouter()
)

func main() {
	const port string = ":10000"

	httpRouter.GET("/", true, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/sources", false, sourceController.GetSources)
	httpRouter.POST("/source", sourceController.AddSource)
	httpRouter.PUT("/source/{id}", sourceController.AddSource)
	httpRouter.GET("/source/{id}", false, sourceController.GetSource)
	httpRouter.DELETE("/source/{id}", sourceController.DeleteSource)

	httpRouter.SERVE(port)

}
