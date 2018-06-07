<template>
	<div>
		<div class="actions">
			<div class="container">
				<div class="row">
					<div class="col-sm-6 col-lg-4 offset-lg-2">
						<div class="search">
							<input class="form-control" type="text" placeholder="Search for a service" v-model="displayOnlyContaining">
							<button class="search_button" value="">
								<svg viewBox="0 0 13 13" width="14" height="14"><title>search</title><path d="M9.448 7.843L13 11.383 11.384 13 7.832 9.448c-.786.501-1.719.797-2.72.797A5.113 5.113 0 0 1 0 5.123 5.116 5.116 0 0 1 5.123 0a5.116 5.116 0 0 1 5.122 5.123 5.077 5.077 0 0 1-.797 2.72zm-4.325.125a2.847 2.847 0 0 0 0-5.691 2.847 2.847 0 0 0 0 5.691z" fill="currentColor" fill-rule="evenodd"></path></svg>
							</button>
						</div>
					</div>
					<div class="col-sm-6 col-lg-6 d-flex align-items-center">
						<div class="form-switch mr-4">
							<input type="checkbox" class="form-switch-input" id="onlyOnline" v-model="displayOnlyOnline">
							<label class="form-switch-label" for="onlyOnline">Only online</label>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div class="container">
			<section>
				<div class="table-responsive">
					<table class="table table-bordered">
						<thead>
							<tr>
								<th scope="col" class="sorting ascending">
									<button>
										Status
										<svg width="8" height="12" viewBox="0 0 8 12" name="sortingArrow">
											<g id="arrow-down" fill="none" fill-rule="evenodd">
												<path id="Shape" fill="currentColor" fill-rule="nonzero" transform="rotate(90 4 5)" d="M5 4h-6v2h6v3l4-4-4-4z">
												</path>
											</g>
										</svg>
									</button>
								</th>
								<th scope="col">
									<button>Organisation</button>
								</th>
								<th scope="col">
									<button>Service</button>
								</th>
								<th scope="col">
									<button>API</button>
								</th>
							</tr>
						</thead>
						<tbody v-if="filteredAndSortedServices && filteredAndSortedServices.length">
							<tr v-for="service of filteredAndSortedServices" v-bind:key="service.organization_name + service.name" v-bind:class="{ 'status-inactive': !service.inway_addresses }">
								<td>
									<svg viewBox="0 0 10 10" width="10" height="10"><title>Oval</title><circle cx="5" cy="14" r="5" transform="translate(0 -9)" fill="currentColor" fill-rule="evenodd"></circle></svg>
								</td>
								<td><span>{{service.organization_name}}</span></td>
								<td><span>{{service.name}}</span></td>
								<td>
									<button type="button" class="btn btn-icon" data-toggle="tooltip" title="" data-original-title="Copy API url">
										<svg viewBox="0 0 32 32" width="32" height="32"><title>Artboard</title><g fill="none" fill-rule="evenodd"><path d="M7.182 7.081a12.454 12.454 0 0 0-3.661 8.84 1.5 1.5 0 1 1-3 0c0-4.084 1.594-8.015 4.54-10.96 6.053-6.054 15.867-6.054 21.92 0l1.06 1.06-6.95 6.95a1.5 1.5 0 1 1-2.12-2.122l4.765-4.765c-4.905-3.858-12.031-3.525-16.554.997z" fill="#FEBF24" fill-rule="nonzero"></path><path d="M24.738 24.839A12.454 12.454 0 0 0 28.4 16a1.5 1.5 0 0 1 3 0c0 4.083-1.594 8.015-4.54 10.96-6.052 6.053-15.867 6.053-21.92 0L3.88 25.9l6.51-6.511a1.5 1.5 0 1 1 2.122 2.121l-4.327 4.327c4.906 3.857 12.032 3.524 16.554-.998z" fill="#CED4DA" fill-rule="nonzero"></path><path d="M14.607 22.839l-1.061 1.06a4 4 0 1 1-5.657-5.656l1.06-1.061a1.5 1.5 0 0 1 2.122-2.121l5.657 5.657a1.5 1.5 0 0 1-2.121 2.12z" fill="#CED4DA"></path><path d="M8.95 17.182a1.5 1.5 0 1 1 2.121-2.121l5.657 5.657a1.5 1.5 0 1 1-2.121 2.12L8.95 17.183z" fill="#E6E8EB"></path><g fill="#FBB301" fill-rule="nonzero"><path d="M16.793 10.793a1 1 0 1 1 1.414 1.414l-2.55 2.55a1 1 0 0 1-1.414-1.415l2.55-2.55zM19.793 13.793a1 1 0 0 1 1.414 1.414l-2.55 2.55a1 1 0 0 1-1.414-1.415l2.55-2.55z"></path></g><path d="M17.182 8.95l1.06-1.06a4 4 0 1 1 5.657 5.656l-1.06 1.06a1.5 1.5 0 0 1-2.121 2.122L15.06 11.07a1.5 1.5 0 0 1 2.121-2.121z" fill="#FEBF24"></path></g></svg>
									</button>
								</td>
							</tr>
						</tbody>
					</table>
				</div>
			</section>
		</div>
	</div>
</template>

<script>
import axios from 'axios'

export default {
	name: 'Directory',
	data () {
		return {
			displayOnlyContaining: '',
			displayOnlyOnline: false,
			sortBy: 'status',
			sortDecending: true,
			services: [],
			errors: []
		}
	},
	created () {
		axios.get(`/api/directory/list-services`)
			.then(response => {
				const { data } = response
				this.services = data.services
			})
			.catch(e => {
				this.errors.push(e)
			})
	},
	computed: {
		filteredAndSortedServices() {
			return this.services.filter(service => {
				if (this.displayOnlyOnline) {
					if (!service.inway_addresses) {
						return false
					}
				}

				if (this.displayOnlyContaining) {
					if (
						!service.name.toLowerCase().includes(this.displayOnlyContaining.toLowerCase()) &&
						!service.organization_name.toLowerCase().includes(this.displayOnlyContaining.toLowerCase())
					) {
						return false
					}
				}

				return true
			})
		}
	}
}
</script>
