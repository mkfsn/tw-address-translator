<script lang="ts">
	import {Container, Alert, Input, Button, Row, Col, Table} from 'sveltestrap'
	import { history } from './stores'

	let input: string;

	async function submit() {
		console.log(input)

		const res = await fetch('https://tw-address-translator.vercel.app/api?address='+input, {
			method: 'GET',
		})

		const result = await res.text()

		console.log(result)

		history.update((obj) => {
			if (!(input in obj)) {
				obj[input] = {created_at: new Date(), count: 0}
			}

			obj[input].updated_at = new Date()
			obj[input].count++
			obj[input].result = result

			return obj
		})
	}
</script>


<svelte:head>
	<title>Taiwan Address Translator</title>
	<meta name="robots" content="noindex nofollow" />
	<meta name="description" content="Taiwan Address Translator">
</svelte:head>

<main>
	<Container>
		<Alert color="primary">
			<h4 class="alert-heading text-capitalize">Taiwan Address Translator</h4>
			<Row>
				<Col md="11">
					<Input
						type="address"
						name="address"
						placeholder="address"
						bind:value={input}
					/>
				</Col>
				<Col md="1">
					<Button block on:click={_ => submit()}>Submit</Button>
				</Col>
			</Row>
		</Alert>
		<Table responsive bordered hover>
			<thead>
				<tr>
					<th>中文地址</th>
					<th>英文地址</th>
					<th>查詢次數</th>
					<th>最後查詢時間</th>
				</tr>
			</thead>
			<tbody>
			{#each Object.entries($history) as [key, record]}
				<tr>
					<td>{key}</td>
					<td>{typeof(record.result) == "string" ? record.result : "UNKNOWN" }</td>
					<td>{record.count}</td>
					<td>{record.updated_at}</td>
				</tr>
			{/each}
			</tbody>
		</Table>
	</Container>

</main>

<style>
	@import "https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css";

	main {
		padding-top: 24px;
	}
</style>
