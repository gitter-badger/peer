var gulp = require("gulp");
var babel = require("gulp-babel");
var browserify = require('gulp-browserify');

gulp.task('default', [
    'js'
]);

gulp.task("js", function () {
    gulp.src("assets/js/app.js")
        .pipe(babel())
        .pipe(browserify())
        .pipe(gulp.dest("public"));
});
