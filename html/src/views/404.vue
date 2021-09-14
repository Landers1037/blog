<template>
    <div style="text-align: center;padding: 5rem 1rem 2rem 1rem">
        <h1>404</h1>
        <p id="not-found"><strong>Page not found</strong></p>
        <p style="margin: 2rem 0">The site configured at this address does not contain the requested file.</p>
        <p style="margin: 2rem 0">十分抱歉, 你请求的页面资源不存在</p>
        <p style="margin: 2rem 0">This is the page generated with <a href="https://github.com/landers1037/blog">Landers1037/blog</a>, try it yourself.</p>
        <p style="margin-top: 4rem;font-weight: 400">Copyright renj.io</p>
        <p style="margin-top: 1rem; cursor: pointer;font-size: 1.2rem"><a href="/">👻</a></p>
    </div>
</template>

<script>
import customData from "../custom/custom";
export default {
    name: "404"
}
</script>

<style scoped>
h1{
    font-family: "Helvetica Neue", Helvetica, Lato, Arial, sans-serif;
    letter-spacing: -1px;
    line-height: 60px;
    font-size: 60px;
    font-weight: 100;
    text-shadow: 0 1px 0 #fff;
}
p{
    color: var(--page-not-found);
}
a {
    color: var(--a-color);
}
a:hover {
    text-decoration: var(--a-color);
}
#not-found {
    margin: 20px 0;
    line-height: 1.6;
}
</style>
<style>
html, body {
    background-color: var(--body-background);
    margin: 0;
    font-family: "Helvetica Neue", Helvetica, Lato, Arial, sans-serif;
}
</style>