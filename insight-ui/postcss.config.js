module.exports = {
    plugins: [
        require('postcss-preset-env')({
            autoprefixer: {
                grid: true,
                flexbox: true,
            },
        }),
    ],
}
