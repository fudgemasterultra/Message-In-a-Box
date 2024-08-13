import { test } from "./lib/test";
(() => {
    const apple = () => {
        console.log('apple');
    }
    apple();
    test();
})();