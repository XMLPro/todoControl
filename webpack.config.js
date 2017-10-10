const path = require('path');

module.exports = {
    entry: './src/js/main.js',
    output: {
        filename: 'bundle.js', 
        path: path.join(__dirname, 'static/build')
    },
    devServer: {
        port: 8001,
        publicPath: '/assets/'
    },
    module: {
        rules: [
            {
                test: /\.vue$/, //loaderの対象ファイル
                exclude: /node_modules/, //loaderの対象外
                use:[{
                    loader: 'vue-loader',
                    options: {
                        loaders: {
                            'js': 'babel-loader!eslint-loader',
                        }
                    }
                }]
            },
            {
                enforce: 'pre',
                test: /\.js$/,
                exclude: /node_modules/,
                loader: 'eslint-loader'
            }
        ]
    }
};
