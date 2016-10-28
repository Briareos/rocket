// Library
var webpack = require('webpack');
var path = require('path');
var HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
    entry: [
        './src/index.jsx'
    ],
    output: {
        path: path.join(__dirname, 'dist'),
        publicPath: '/',
        filename: 'application.js'
    },
    resolve: {
        extensions: ['', '.js', '.jsx']
    },
    module: {
        loaders: [
            {
                test: /\.jsx?$/,
                exclude: /(node_modules|public)/,
                loader: 'babel',
                query: {
                    presets: ['es2015', 'react']
                }
            }
        ]
    },
    devServer: {
        contentBase: "./src",
        noInfo: true,
        hot: true,
        inline: true,
        historyApiFallback: true,
        proxy: {
            "/api": {
                "target": {
                    "host": "api.rocket.dev",
                    "protocol": 'http:',
                    "port": 80
                },
                ignorePath: true,
                changeOrigin: true,
                secure: false
            }
        },
        port: '4000',
        host: '127.0.0.1'
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: './src/index.html'
        })
    ]
};