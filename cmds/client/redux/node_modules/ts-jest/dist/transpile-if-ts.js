"use strict";
var tsc = require("typescript");
var utils_1 = require("./utils");
function transpileIfTypescript(path, contents) {
    if (path && (path.endsWith('.tsx') || path.endsWith('.ts'))) {
        var transpiled = tsc.transpileModule(contents, {
            compilerOptions: utils_1.getTSConfig({ __TS_CONFIG__: global['__TS_CONFIG__'] }, true),
            fileName: path
        });
        return transpiled.outputText;
    }
    return contents;
}
exports.transpileIfTypescript = transpileIfTypescript;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoidHJhbnNwaWxlLWlmLXRzLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiLi4vc3JjL3RyYW5zcGlsZS1pZi10cy50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiO0FBQUEsZ0NBQWtDO0FBQ2xDLGlDQUFzQztBQUV0QywrQkFBc0MsSUFBSSxFQUFFLFFBQVE7SUFDbEQsRUFBRSxDQUFDLENBQUMsSUFBSSxJQUFJLENBQUMsSUFBSSxDQUFDLFFBQVEsQ0FBQyxNQUFNLENBQUMsSUFBSSxJQUFJLENBQUMsUUFBUSxDQUFDLEtBQUssQ0FBQyxDQUFDLENBQUMsQ0FBQyxDQUFDO1FBRTVELElBQUksVUFBVSxHQUFHLEdBQUcsQ0FBQyxlQUFlLENBQUMsUUFBUSxFQUFFO1lBQzdDLGVBQWUsRUFBRSxtQkFBVyxDQUFDLEVBQUUsYUFBYSxFQUFFLE1BQU0sQ0FBQyxlQUFlLENBQUMsRUFBRSxFQUFFLElBQUksQ0FBQztZQUM5RSxRQUFRLEVBQUUsSUFBSTtTQUNmLENBQUMsQ0FBQztRQUVILE1BQU0sQ0FBQyxVQUFVLENBQUMsVUFBVSxDQUFDO0lBQy9CLENBQUM7SUFDRCxNQUFNLENBQUMsUUFBUSxDQUFDO0FBQ2xCLENBQUM7QUFYRCxzREFXQyJ9