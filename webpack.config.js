var path = require('path');
module.exports = {
    output: {
        filename: '[name].js',
        sourceMapFilename : '../map/ts-map/[file].map'
    },
    resolve: {
        modules:[path.join(__dirname,'node_modules')],
        extensions:['*', '.ts', '.webpack.js', '.web.js', '.js']
    },
    module: {
        loaders: [
            { test: /\.ts$/, loader: 'ts-loader' }
        ]
    },
    devtool: "source-map",
}
