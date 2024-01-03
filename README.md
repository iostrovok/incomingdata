# The package provides a simple object to validate incoming data

See ./types.go to support types.

## Example:

```shell
import(
  "api"
  
  inc "github.com/iostrovok/incomingdata.go"
)
  ...
  
func (g *grpcserver) GetByIds(in *api.Parameters) (*api.result, error) {
	err := inc.New().
		Add("list of ids", in.GetIds(), inc.ListInt32P).
		Add("time of request", in.GetTime(), inc.Int64P).
		First()
		
	if err != nil {
        return nil, err
	}
	
	...
```
