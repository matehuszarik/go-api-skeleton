## GO API skeleton

Sample project for bootstrapping HTTP services.

### Setup
1. Make sure [golang/dep](https://github.com/golang/dep) is installed
2. Run `dep init` to initialize dependency management
3. Create `.env` file with the following environment variables
```
API_KEYS=a,b,c
PORT=1234
```
4. Rename the imports by running `find . --type f --exec sed -i '' -e 's,github.com/matehuszarik/go-api-skeleton,name/of/your/path,g' {} +` on macOS