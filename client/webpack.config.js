// Library
var webpack = require('webpack');
var path = require('path');
var HtmlWebpackPlugin = require('html-webpack-plugin');

// Local
var loaders = require('./webpack.loaders');

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
        port: '4000',
        host: '127.0.0.1'
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: './src/index.html'
        })
    ]
};