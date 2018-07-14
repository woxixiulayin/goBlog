import React from 'react'
import Posts from 'src/views/Posts'

const posts = {
    label: 'posts',
    zH: '文章',
    path: '/posts',
    component: Posts
}

const about = {
    label: 'about',
    zH: '关于',
    path: '/about',
    component: () => <div>关于我的info</div>
}

const navMap = {
    posts,
    about
}

// 提供数组数据，便于map语法显示
const navList = Object.values(navMap)

export {
    navMap,
    navList
}