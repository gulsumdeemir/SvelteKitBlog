export async function createPost(post: {
	postTitle: string;
	postDescription: string;
	categoryID: number;
	tags: number[];
}) {
	const token = localStorage.getItem('token');

	if (!token) {
		alert('Giriş yapmalısınız.');
		window.location.href = '/login';
		return;
	}

	console.log('Giriş Tokenı:', token);

	console.log('Gönderilen Yazı:', post);
	console.log('Başlıklar:', {
		'Content-Type': 'application/json',
		Authorization: `Bearer ${token}`
	});

	const response = await fetch('http://localhost:3001/posts', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${token}`
		},
		body: JSON.stringify(post)
	});

	if (!response.ok) {
		const errorData = await response.json().catch(() => ({ message: 'Bilinmeyen hata' }));
		console.error('Sunucu Hatası:', errorData);
		throw new Error(`Yazı oluşturma isteği başarısız oldu: ${errorData.message}`);
	}

	return await response.json();
}

export async function getCategories() {
	const response = await fetch('http://localhost:3001/categories');
	if (!response.ok) {
		throw new Error('Kategorileri alırken hata oluştu.');
	}
	return await response.json();
}

export async function getTags() {
	const response = await fetch('http://localhost:3001/tags');
	if (!response.ok) {
		throw new Error('Etiketleri alırken hata oluştu.');
	}
	return await response.json();
}
