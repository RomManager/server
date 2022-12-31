module.exports = {
  /**
   * This is the main entry point for this application, it's the first file
   * that runs in the main process.
   */
  entry: ["./src/main/app.ts"],
  module: {
    rules: require("./webpack.rules"),
  },
  resolve: {
    extensions: [".js", ".ts", ".jsx", ".tsx", ".css", ".json"],
    alias: require("./webpack.aliases"),
  },
  stats: "minimal",
};
