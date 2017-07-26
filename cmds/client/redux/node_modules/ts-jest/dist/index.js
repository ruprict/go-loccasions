"use strict";
var default_retrieve_file_handler_1 = require("./default-retrieve-file-handler");
var sourceMapSupport = require("source-map-support");
function install() {
    var options = {};
    options.retrieveFile = default_retrieve_file_handler_1.defaultRetrieveFileHandler;
    options.emptyCacheBetweenOperations = true;
    options['environment'] = 'node';
    return sourceMapSupport.install(options);
}
exports.install = install;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiaW5kZXguanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvaW5kZXgudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IjtBQUFBLGlGQUE2RTtBQUM3RSxxREFBdUQ7QUFFdkQ7SUFDRSxJQUFJLE9BQU8sR0FBNkIsRUFBRSxDQUFDO0lBQzNDLE9BQU8sQ0FBQyxZQUFZLEdBQUcsMERBQTBCLENBQUM7SUFDbEQsT0FBTyxDQUFDLDJCQUEyQixHQUFHLElBQUksQ0FBQztJQUMzQyxPQUFPLENBQUMsYUFBYSxDQUFDLEdBQUcsTUFBTSxDQUFDO0lBRWhDLE1BQU0sQ0FBQyxnQkFBZ0IsQ0FBQyxPQUFPLENBQUMsT0FBTyxDQUFDLENBQUM7QUFDM0MsQ0FBQztBQVBELDBCQU9DIn0=