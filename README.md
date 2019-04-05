# cryptcfg
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
