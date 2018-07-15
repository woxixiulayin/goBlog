/**@flow */
import Rx from 'rxjs'
import { EntityModule } from 'src/lib/Module'

const posts = new EntityModule({
    state: {}
})

posts.$$state.subscribe(posts => console.log(posts))
posts.save([{
    id: 1,
    title: '123',
}])
posts.remove([1])
window.posts = posts