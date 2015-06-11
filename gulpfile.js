var gulp = require("gulp");
var babel = require("gulp-babel");
var browserify = require('gulp-browserify');

gulp.task('default', ['js'], function () {
    return gulp.watch('app/assets/js/**', ['js']);
});

gulp.task("js", function () {
    gulp.src("app/assets/js/app.js")
        .pipe(babel())
        .pipe(browserify())
        .pipe(gulp.dest("public"));
});
