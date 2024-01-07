# assets

Dead simple Vite + Go integration.

## Usage

1. Install and configure Vite.

Your `vite.config.js` should look like this:

```js
import { defineConfig } from "vite"

export default defineConfig({
  build: {
    manifest: true,
    outDir: "web/dist",
    rollupOptions: {
      input: ["web/index.js", "web/index.css"],
    },
  }
})
```

1. Create a `Vite` struct.

For production, the config should look like this:

```go
vite := assets.Vite{
    ManifestPath: "web/dist/.vite/manifest.json",
    Mode:         assets.ModeProduction,
    StaticURL:    "/static",
}
```

For development, the config should look like this:

```go
vite := assets.Vite{
    Mode:         assets.ModeDevelopment,
    DevServerURL: "http://localhost:5173",
}
```

2. Before parsing templates, add a the `Resolve` method to the `FuncMap`.

```go
tm := template.New("page").Funcs(template.FuncMap{
    "vite": vite.Resolve,
})
```

3. Inside your templates, use the `vite` function to resolve assets.

```html
{{vite "web/index.js" "web/index.css" }}
```

4. Profit!
