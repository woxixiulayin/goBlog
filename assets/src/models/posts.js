// @flow
import ajax from 'src/lib/ajax'
import entityUtils from 'src/lib/entityUtils'
import { createModel } from '@rematch/core'
import * as postsOperations from 'src/ajaxOperations/posts'

export type typePostEntity = {
    id: string,
    title: string,
    tags: Array<string>,
    pv: number,
    user_id: string,
    comments: Array<string>,
}

type typePostsById = {
    [id: string]: typePostEntity
}

type typeState = {
    postsById: typePostsById
}

const posts = createModel({
    state: {
        postsById: {}
    },
    reducers: {
        setPosts(state, posts: Array<typePostEntity>): typePostsById {
            return {
                ...state,
                postsById: entityUtils.setEntities(posts)
            }
        },
        savePosts(state, posts: Array<typePostEntity> | typePostEntity): typePostsById {
            return {
                ...state,
                postsById: entityUtils.saveEntities(state, posts),
            }
        },
        deletePosts(state, postsIds: Array<string> | string): typePostsById {
            return {
                ...state,
                postsById: entityUtils.deleteEntities(state, postsIds)
            }
        }
    },
    effects: dispatch => ({
        async fetchPosts(page = 0) {
            const posts = await postsOperations.fetchPosts({ page })
            dispatch.posts.savePosts(posts)
            return posts
        }
    })
})

export default posts