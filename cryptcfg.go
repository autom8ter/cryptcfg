package cryptcfg

import (
	"errors"
	"github.com/autom8ter/gocrypt"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Crypt struct {
	v *viper.Viper
}

func New() *Crypt {
	v := viper.GetViper()
	if v == nil {
		v = viper.New()
	}
	return &Crypt{
		v: v,
	}
}

func (c *Crypt) Viper() *viper.Viper {
	return c.v
}

func (c *Crypt) ReadInEncrypted(path, key string) error {
	if path == "" {
		return errors.New("please provide a valid file path")
	}
	var newPath string
	if filepath.Ext(path) == "enc" {
		newPath = strings.TrimSuffix(path, ".enc")
	} else {
		newPath = path
	}
	g := gocrypt.NewGoCrypt()
	err := g.DecryptFiles(path, key, 0755)
	if err != nil {
		return err
	}

	bits, err := ioutil.ReadFile(newPath)
	if err != nil {
		return err
	}
	if err := viper.ReadConfig(strings.NewReader(string(strings.TrimSpace(string(bits))))); err != nil {
		return err
	}
	err = g.EncryptFiles(newPath, key, 0755)
	if err != nil {
		return err
	}
	return nil
}
