const stats = {
    // fallback value for stats options when an option is not defined (has precedence over local webpack defaults)
    all: undefined,
    // Add asset Information
    assets: true,
    // Sort assets by a field
    // You can reverse the sort with `!field`.
    // assetsSort: "!field",
    // Add build date and time information
    builtAt: false,
    // Add information about cached (not built) modules
    cached: false,
    // Show cached assets (setting this to `false` only shows emitted files)
    cachedAssets: false,
    // Add children information
    children: false,
    // `webpack --colors` equivalent
    colors: true,
    // Display the entry points with the corresponding bundles
    entrypoints: false,
    // Add --env information
    env: true,
    // Add errors
    errors: true,
    // Add details to errors (like resolving log)
    errorDetails: true,
    // Add the hash of the compilation
    hash: false,
    // Set the maximum number of modules to be shown
    // maxModules: 15,
    // Add built modules information
    modules: false,
    // Add the source code of modules
    source: false,
}

module.exports = { stats }
