var browserify = require('browserify');
var gulp = require('gulp');
var source = require('vinyl-source-stream');
var buffer = require('vinyl-buffer');
var gutil = require('gulp-util');
var babelify = require("babelify");

gulp.task('default', [
    'js'
]);

gulp.task('watch', ['default'], function () {
    return gulp.watch('assets/js/**', ['js']);
});

gulp.task('js', function () {
    var b = browserify({
        entries: 'assets/js/app.js',
        debug: true
    });

    return b
        .transform(babelify)
        .bundle()
        .pipe(source('app.js'))
        .pipe(buffer())
        .on('error', gutil.log)
        .pipe(gulp.dest('./public/'));
});
