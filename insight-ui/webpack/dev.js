// const webpack = require('webpack');
const path = require('path')

const HtmlWebpackPlugin = require('html-webpack-plugin')
const CopyWebpackPlugin = require('copy-webpack-plugin')

const MiniCssExtractPlugin = require('mini-css-extract-plugin')

// load stats configuration
const { stats } = require('./stats')
// construct dist folder location
// note: we go up from webpack folder into dist
const dist = path.resolve(__dirname, '../dist')
// console.log("dist...", dist);

// load proxy settings
let p = require('./proxy')

module.exports = (env) => {
    return {
        mode: 'development',
        entry: {
            index: './src/index.js',
        },
        output: {
            filename: '[name].js',
            chunkFilename: '[name].js',
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
                                sourceMap: true,
                            },
                        },
                        {
                            loader: 'postcss-loader',
                            options: {
                                sourceMap: true,
                            },
                        },
                        {
                            loader: 'sass-loader',
                            options: {
                                sourceMap: true,
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
            // copy index html
            // https://webpack.js.org/plugins/html-webpack-plugin/
            new HtmlWebpackPlugin({
                filename: 'index.html',
                template: './src/index.html',
                inject: true,
                favicon: './static/favicon.ico',
            }),
            // old extract text plugin to extract css
            // new ExtractTextPlugin('[name].css')
            new MiniCssExtractPlugin({
                // Options similar to webpackOptions.output
                // both options are optional
                filename: '[name].css',
                chunkFilename: '[id].css',
            }),
            // copy assets
            // https://webpack.js.org/plugins/copy-webpack-plugin/
            new CopyWebpackPlugin([
                // copy all files from static dir to root
                // note: when no files folder is not copied!
                './static/',
            ]),
        ],
        /**
         * Display stats, see link below for complete list
         * https://webpack.js.org/configuration/stats/#stats
         */
        stats: stats,
        /**
         * Webpack dev server setup
         */
        devtool: 'source-map',
        devServer: {
            port: 3000,
            stats: stats,
            compress: true,
            // https: true,
            // route rewrites
            historyApiFallback: true,
            proxy: p(env),
        },
    }
}
