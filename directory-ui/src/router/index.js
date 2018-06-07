import Vue from 'vue'
import Router from 'vue-router'
import Directory from '@/Directory'

Vue.use(Router)

export default new Router({
	mode: 'history',
	routes: [
		{
			path: '/',
			name: 'Directory',
			component: Directory
		}
	]
})
