var gulp = require("gulp");
var babel = require("gulp-babel");
var browserify = require('gulp-browserify');

gulp.task('default', ['js'], function () {
    return gulp.watch('assets/js/**', ['js']);
});

gulp.task("js", function () {
    gulp.src("assets/js/app.js")
        .pipe(babel())
        .pipe(browserify())
        .pipe(gulp.dest("public"));
});
