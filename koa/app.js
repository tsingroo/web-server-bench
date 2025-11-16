const Koa = require('koa');
const Router = require('@koa/router');

const app = new Koa();
const router = new Router();

// 定义路由
router.get('/ping', (ctx) => {
  ctx.body = 'pong';
});

router.get('/hello/:name', (ctx) => {
  const { name } = ctx.params;
  ctx.body = `my name is ${name}`;
});


// 使用路由中间件
app.use(router.routes());

app.listen(8080);
