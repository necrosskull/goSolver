<script xmlns:svelte="http://www.w3.org/1999/html">
	$: searchQuery = '';
	let searchResults = [];
	let timeout;

	async function search() {
		clearTimeout(timeout);

		if (searchQuery.trim().length < 2 || searchQuery === '') {
			searchResults = [];
			return;
		}

		timeout = setTimeout(async () => {
			const response = await fetch(`http://localhost:8082/api/v1/search/${searchQuery}`, {
				method: 'GET'
			});

			searchResults = await response.json();
		}, 300);
	}
</script>

<svelte:head>
	<title>Поиск по ответам</title>
	<meta name="description" content="Поиск по ответам" />
</svelte:head>

<main class="">
	<div
		class=" flex flex-col items-center min-h-[calc(100dvh)] max-w-[calc(100dvw)] bg-neutral-900 text-white"
	>
		<div class="min-h-[calc(100dvh/1.8)] flex flex-col items-center justify-end w-full">
			<h1 class="text-4xl font-bold mt-8 mb-4">Поиск по ответам</h1>
			<input
				class="rounded-xl p-4 w-11/12 bg-neutral-800 text-white outline-none lg:w-1/2 mb-4"
				type="text"
				bind:value={searchQuery}
				on:input={search}
				placeholder="Поиск..."
			/>
		</div>

		<ul class="w-11/12 space-y-4 mb-8 lg:w-1/2">
			{#each searchResults as result (result._id.$oid)}
				{#if result.answers.length > 0 && !result.question.includes('Вы можете помочь развитию проекта')}
					<li class="p-4 bg-neutral-800 rounded-xl">
						<h3 class="mb-4 overflow-x-auto">{result.question}</h3>

						{#if result.answers.length > 0}
							<ul class="space-y-2">
								{#each result.answers.sort((a, b) => b.users.length - a.users.length) as answers}
									{#if answers.users.length > 0 || answers.correct.length > 0}
										{#if answers.correct.length > 0}
											<li
												class="flex items-center p-4 bg-green-900 rounded-xl justify-between overflow-x-auto"
												title="Ответ отмеченный как верный"
											>
												<span class="">{answers.answer}</span>
												<span class="text-neutral-300" title="Количество людей выбравших этот ответ"
													>({answers.users.length})</span
												>
											</li>
										{:else if answers.not_correct.length > 0}
											<li
												class="flex items-center p-4 bg-red-900 rounded-xl justify-between overflow-x-auto"
												title="Ответ отмеченный как неверный"
											>
												<span class="">{answers.answer}</span>
												<span class="text-neutral-300" title="Количество людей выбравших этот ответ"
													>({answers.users.length})</span
												>
											</li>
										{:else}
											<li
												class="flex items-center p-4 bg-neutral-700 rounded-xl justify-between overflow-x-auto"
												title="Ответ выбранный хотябы одним пользователем"
											>
												<span class="">{answers.answer}</span>
												<span class="text-neutral-300" title="Количество людей выбравших этот ответ"
													>({answers.users.length})</span
												>
											</li>
										{/if}
									{/if}
								{/each}
							</ul>
						{/if}
					</li>
				{/if}
			{/each}
		</ul>
	</div>
</main>
