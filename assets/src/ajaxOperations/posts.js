import ajax from 'src/lib/ajax'

export const fetchPosts = ({ page = 0 }) => ajax.get(`/posts?user=1&page=${page}&page_size=10`)