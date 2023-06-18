const go = new window.Go();
WebAssembly.instantiateStreaming(fetch("analyze.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});