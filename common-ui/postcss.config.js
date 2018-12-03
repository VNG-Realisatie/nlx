module.exports = () => {
    return {
        plugins:[
            require('postcss-preset-env')({
                autoprefixer: {
                    grid: true,
                    flexbox: true
                }
            })
        ]
    }
}
