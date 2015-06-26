var browserify = require('browserify');
var gulp = require('gulp');
var source = require('vinyl-source-stream');
var buffer = require('vinyl-buffer');
var gutil = require('gulp-util');
var babelify = require("babelify");
var sass = require('gulp-sass');

gulp.task('default', [
    'js',
    'css'
]);

gulp.task('watch', ['default'], function () {
    gulp.watch('assets/js/**/*.js', ['js']);
    gulp.watch('assets/css/**/*.scss', ['css']);
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

gulp.task('css', function () {
    gulp.src('assets/css/**/*.scss')
        .pipe(sass().on('error', sass.logError))
        .pipe(gulp.dest('public'));
});
