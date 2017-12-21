//plug-in
var gulp = require('gulp');
var pug = require('gulp-pug');
var sass = require('gulp-sass');
var ts = require('gulp-typescript');
var sourcemaps = require('gulp-sourcemaps');
var webpack = require('webpack-stream');
var named = require('vinyl-named');
var plumber = require('gulp-plumber');
 

gulp.task('pug', () => {
    gulp.src(['./src/pug/**/*.pug', '!./src/pug/**/_*.pug'], { base: 'src/pug' })
        .pipe(plumber())
        .pipe(pug({
            pretty: true
        }))
        .pipe(gulp.dest('./views/'));
});


gulp.task('sass', function() {
    gulp.src(['./src/sass/**/*.sass', '!./src/sass/**/_*.sass'],{ base: 'src/sass' })
        .pipe(plumber())
        .pipe(sourcemaps.init())
        .pipe(sass())
        .pipe(sourcemaps.write('../map/sass-map'))
        .pipe(gulp.dest('./public/css/'));
});


var webpackConfig = require('./webpack.config.js');
 
gulp.task('ts', function () {
    gulp.src(['./src/ts/**/*.ts', '!./src/ts/**/_*.ts','!./node_modules/**'],{base:'src/ts'})
    .pipe(plumber())
    .pipe(named(function(file) {
      return file.relative.replace(/\.[^\.]+$/, '');
    }))
    .pipe(webpack(webpackConfig))
    .pipe(gulp.dest('./public/js/'));
});


gulp.task('default', ['pug','sass','ts'], function() {
    gulp.watch(['./src/pug/**/*.pug'], ['pug']);
    gulp.watch(['./src/sass/**/*.sass'], ['sass']);
    gulp.watch(['./src/ts/**/*.ts'], ['ts']);
})
