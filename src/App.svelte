<script lang="ts">
	import {Container, Alert, Input, Button, Row, Col} from 'sveltestrap';
</script>

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
					/>
				</Col>
				<Col md="1">
					<Button block>Submit</Button>
				</Col>
			</Row>
		</Alert>
	</Container>

</main>

<style>
	@import "https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css";

	main {
		padding-top: 24px;
	}
</style>
