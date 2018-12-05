// const webpack = require('webpack');
const path = require('path')

const HtmlWebpackPlugin = require('html-webpack-plugin')
const CopyWebpackPlugin = require('copy-webpack-plugin')
const CleanWebpackPlugin = require('clean-webpack-plugin')
// const BundleAnalyzerPlugin = require('webpack-bundle-analyzer')
//     .BundleAnalyzerPlugin
/*
 * We've enabled UglifyJSPlugin for you! This minifies your app
 * in order to load faster and run less javascript.
 * https://github.com/webpack-contrib/uglifyjs-webpack-plugin
 */
const UglifyJSPlugin = require('uglifyjs-webpack-plugin')

// const ExtractTextPlugin = require('extract-text-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const OptimizeCSSAssetsPlugin = require('optimize-css-assets-webpack-plugin')

const dist = path.resolve(__dirname, '../build')

module.exports = {
    mode: 'production',
    entry: {
        index: './src/index.js',
    },
    output: {
        filename: '[name].[chunkhash].js',
        chunkFilename: '[name].[chunkhash].js',
        path: dist,
        publicPath: '/',
    },

    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                loader: 'babel-loader',
            },
            {
                test: /\.(scss|css)$/,
                use: [
                    // extract css into separate file
                    MiniCssExtractPlugin.loader,
                    {
                        loader: 'css-loader',
                        options: {
                            sourceMap: false,
                        },
                    },
                    {
                        loader: 'postcss-loader',
                        options: {
                            sourceMap: false,
                        },
                    },
                    {
                        loader: 'sass-loader',
                        options: {
                            sourceMap: false,
                        },
                    },
                ],
            },
            {
                test: /\.(png|jpg|gif|svg)$/i,
                use: [
                    {
                        loader: 'url-loader',
                        options: {
                            limit: 2048,
                            name: 'img/[name].[ext]',
                        },
                    },
                ],
            },
            {
                test: /\.(woff(2)?|ttf|eot|svg)$/i,
                use: [
                    {
                        loader: 'url-loader',
                        options: {
                            limit: 1024,
                            name: 'font/[name].[ext]',
                        },
                    },
                ],
            },
        ],
    },

    plugins: [
        // remove all files from dist folder on each build
        new CleanWebpackPlugin(['build'], {
            root: path.resolve(__dirname, '..'),
        }),
        // copy index html
        // https://webpack.js.org/plugins/html-webpack-plugin/
        new HtmlWebpackPlugin({
            template: './src/index.html',
        }),
        // extract css to separate file
        // https://webpack.js.org/plugins/mini-css-extract-plugin/
        new MiniCssExtractPlugin({
            // Options similar to webpackOptions.output
            // both options are optional
            filename: '[name].[chunkhash].css',
            chunkFilename: '[id].[chunkhash].css',
        }),
        // copy assets
        // https://webpack.js.org/plugins/copy-webpack-plugin/
        new CopyWebpackPlugin([
            // copy all files from assets dir to root
            './static/',
        ]),
        // uglify js
        new UglifyJSPlugin(),
        // optimize css
        new OptimizeCSSAssetsPlugin(),
        /* further investigation needed
			https://www.npmjs.com/package/webpack-bundle-analyzer
        new BundleAnalyzerPlugin({
            generateStatsFile: true,
        }), */
    ],
}
