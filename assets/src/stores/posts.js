/**@flow */
import Rx from 'rxjs'
import Module from 'src/lib/Module'

const posts = new Module({
    state: {}
})

posts.$$state.subscribe(posts => console.log(posts))
posts.save([{
    id: 1,
    title: '123',
}])
posts.remove([1])
window.posts = posts