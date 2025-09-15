fn main() {
    // Build script for Node.js bindings
    #[cfg(not(target_arch = "wasm32"))]
    {
        napi_build::setup();
    }
}