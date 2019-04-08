# cryptcfg

a golang library

cryptcfg uses a key to decrypt a file thats been encrypted with the gocrypt cli tool at runtime, read in the config, and store the config in memory without ever decrypting the file on disk.

--
    import "github.com/autom8ter/cryptcfg"


## Usage

#### type Crypt

```go
type Crypt struct {
}
```


#### func  New

```go
func New() *Crypt
```

#### func (*Crypt) ReadInEncrypted

```go
func (c *Crypt) ReadInEncrypted(path, key string) error
```

#### func (*Crypt) Viper

```go
func (c *Crypt) Viper() *viper.Viper
```
