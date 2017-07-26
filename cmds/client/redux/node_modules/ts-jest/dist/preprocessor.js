"use strict";
var tsc = require("typescript");
var nodepath = require("path");
var utils_1 = require("./utils");
var glob = require('glob-all');
var getPackageRoot = require('jest-util').getPackageRoot;
var root = getPackageRoot();
var _a = utils_1.getJestConfig(root), testRegex = _a.testRegex, collectCoverage = _a.collectCoverage, coverageDirectory = _a.coverageDirectory, coverageReporters = _a.coverageReporters, collectCoverageFrom = _a.collectCoverageFrom, testResultsProcessor = _a.testResultsProcessor;
if (testResultsProcessor) {
    global.__ts_coverage__cache__ = {};
    global.__ts_coverage__cache__.sourceCache = {};
    global.__ts_coverage__cache__.coverageConfig = { collectCoverage: collectCoverage, coverageDirectory: coverageDirectory, coverageReporters: coverageReporters };
    global.__ts_coverage__cache__.coverageCollectFiles =
        collectCoverage &&
            testResultsProcessor &&
            collectCoverageFrom &&
            collectCoverageFrom.length ?
            glob.sync(collectCoverageFrom).map(function (x) { return nodepath.resolve(root, x); }) : [];
}
function process(src, path, config) {
    if (path.endsWith('.ts') || path.endsWith('.tsx')) {
        var transpiled = tsc.transpileModule(src, {
            compilerOptions: utils_1.getTSConfig(config.globals, collectCoverage),
            fileName: path
        });
        if (global.__ts_coverage__cache__) {
            if (!testRegex || !path.match(testRegex)) {
                global.__ts_coverage__cache__.sourceCache[path] = transpiled.outputText;
            }
        }
        var start = transpiled.outputText.length > 12 ? transpiled.outputText.substr(1, 10) : '';
        var modified = start === 'use strict'
            ? "'use strict';require('ts-jest').install();" + transpiled.outputText
            : "require('ts-jest').install();" + transpiled.outputText;
        return modified;
    }
    return src;
}
exports.process = process;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoicHJlcHJvY2Vzc29yLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiLi4vc3JjL3ByZXByb2Nlc3Nvci50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiO0FBQUEsZ0NBQWtDO0FBQ2xDLCtCQUFpQztBQUNqQyxpQ0FBcUQ7QUFFckQsSUFBTSxJQUFJLEdBQUcsT0FBTyxDQUFDLFVBQVUsQ0FBQyxDQUFDO0FBQ2pDLElBQU0sY0FBYyxHQUFHLE9BQU8sQ0FBQyxXQUFXLENBQUMsQ0FBQyxjQUFjLENBQUM7QUFJM0QsSUFBTSxJQUFJLEdBQUcsY0FBYyxFQUFFLENBQUM7QUFDeEIsSUFBQSxnQ0FPaUIsRUFObkIsd0JBQVMsRUFDVCxvQ0FBZSxFQUNmLHdDQUFpQixFQUNqQix3Q0FBaUIsRUFDakIsNENBQW1CLEVBQ25CLDhDQUFvQixDQUNBO0FBR3hCLEVBQUUsQ0FBQyxDQUFDLG9CQUFvQixDQUFDLENBQUMsQ0FBQztJQUN2QixNQUFNLENBQUMsc0JBQXNCLEdBQUcsRUFBRSxDQUFDO0lBQ25DLE1BQU0sQ0FBQyxzQkFBc0IsQ0FBQyxXQUFXLEdBQUcsRUFBRSxDQUFDO0lBQy9DLE1BQU0sQ0FBQyxzQkFBc0IsQ0FBQyxjQUFjLEdBQUcsRUFBRSxlQUFlLGlCQUFBLEVBQUUsaUJBQWlCLG1CQUFBLEVBQUUsaUJBQWlCLG1CQUFBLEVBQUUsQ0FBQztJQUN6RyxNQUFNLENBQUMsc0JBQXNCLENBQUMsb0JBQW9CO1FBQzlDLGVBQWU7WUFDWCxvQkFBb0I7WUFDcEIsbUJBQW1CO1lBQ25CLG1CQUFtQixDQUFDLE1BQU07WUFDMUIsSUFBSSxDQUFDLElBQUksQ0FBQyxtQkFBbUIsQ0FBQyxDQUFDLEdBQUcsQ0FBQyxVQUFBLENBQUMsSUFBSSxPQUFBLFFBQVEsQ0FBQyxPQUFPLENBQUMsSUFBSSxFQUFFLENBQUMsQ0FBQyxFQUF6QixDQUF5QixDQUFDLEdBQUcsRUFBRSxDQUFDO0FBQ3BGLENBQUM7QUFFRCxpQkFBd0IsR0FBRyxFQUFFLElBQUksRUFBRSxNQUFNO0lBQ3JDLEVBQUUsQ0FBQyxDQUFDLElBQUksQ0FBQyxRQUFRLENBQUMsS0FBSyxDQUFDLElBQUksSUFBSSxDQUFDLFFBQVEsQ0FBQyxNQUFNLENBQUMsQ0FBQyxDQUFDLENBQUM7UUFDaEQsSUFBTSxVQUFVLEdBQUcsR0FBRyxDQUFDLGVBQWUsQ0FDbEMsR0FBRyxFQUNIO1lBQ0ksZUFBZSxFQUFFLG1CQUFXLENBQUMsTUFBTSxDQUFDLE9BQU8sRUFBRSxlQUFlLENBQUM7WUFDN0QsUUFBUSxFQUFFLElBQUk7U0FDakIsQ0FBQyxDQUFDO1FBR1AsRUFBRSxDQUFDLENBQUMsTUFBTSxDQUFDLHNCQUFzQixDQUFDLENBQUMsQ0FBQztZQUNoQyxFQUFFLENBQUMsQ0FBQyxDQUFDLFNBQVMsSUFBSSxDQUFDLElBQUksQ0FBQyxLQUFLLENBQUMsU0FBUyxDQUFDLENBQUMsQ0FBQyxDQUFDO2dCQUN2QyxNQUFNLENBQUMsc0JBQXNCLENBQUMsV0FBVyxDQUFDLElBQUksQ0FBQyxHQUFHLFVBQVUsQ0FBQyxVQUFVLENBQUM7WUFDNUUsQ0FBQztRQUNMLENBQUM7UUFFRCxJQUFNLEtBQUssR0FBRyxVQUFVLENBQUMsVUFBVSxDQUFDLE1BQU0sR0FBRyxFQUFFLEdBQUcsVUFBVSxDQUFDLFVBQVUsQ0FBQyxNQUFNLENBQUMsQ0FBQyxFQUFFLEVBQUUsQ0FBQyxHQUFHLEVBQUUsQ0FBQztRQUUzRixJQUFNLFFBQVEsR0FBRyxLQUFLLEtBQUssWUFBWTtjQUNqQywrQ0FBNkMsVUFBVSxDQUFDLFVBQVk7Y0FDcEUsa0NBQWdDLFVBQVUsQ0FBQyxVQUFZLENBQUM7UUFFOUQsTUFBTSxDQUFDLFFBQVEsQ0FBQztJQUNwQixDQUFDO0lBRUQsTUFBTSxDQUFDLEdBQUcsQ0FBQztBQUNmLENBQUM7QUExQkQsMEJBMEJDIn0=