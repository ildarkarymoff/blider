package builder

import (
	"fmt"
	"github.com/ildarkarymoff/blider/config"
	"github.com/ildarkarymoff/blider/repository"
	"os/exec"
	"path/filepath"
)

type GnomeCmdBuilder struct {
	config *config.Config
}

func (b *GnomeCmdBuilder) Init(config *config.Config) {
	b.config = config
}

func (b *GnomeCmdBuilder) Build(wallpaper *repository.Wallpaper) *exec.Cmd {
	imgPath := filepath.Join(b.config.LocalStoragePath, wallpaper.Filename)

	return exec.Command(
		"gsettings",
		"set",
		"org.gnome.desktop.background",
		"picture-uri",
		fmt.Sprintf("file:///%s", imgPath),
	)
}
