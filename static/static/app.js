document.addEventListener('DOMContentLoaded', () => {
	const container = document.getElementById('container')
	const signUpBtn = document.getElementById('signUp')
	const signInBtn = document.getElementById('signIn')
	const mobileSignIn = document.getElementById('mobileSwitchToSignIn')
	const mobileSignUp = document.getElementById('mobileSwitchToSignUp')

	// Переключение панелей
	const togglePanel = isSignUp => {
		container.classList.toggle('right-panel-active', isSignUp)
		updateMobileSwitcher()
	}

	// Обновление мобильных ссылок
	const updateMobileSwitcher = () => {
		const isActive = container.classList.contains('right-panel-active')
		mobileSignIn.style.display = isActive ? 'inline-block' : 'none'
		mobileSignUp.style.display = isActive ? 'none' : 'inline-block'
	}

	// Обработчики событий
	signUpBtn.addEventListener('click', () => togglePanel(true))
	signInBtn.addEventListener('click', () => togglePanel(false))
	mobileSignIn.addEventListener('click', e => {
		e.preventDefault()
		togglePanel(false)
	})
	mobileSignUp.addEventListener('click', e => {
		e.preventDefault()
		togglePanel(true)
	})

	// Отправка форм
	const handleForm = async (form, url) => {
		const formData = {
			name: form.querySelector('[name="name"]')?.value,
			email: form.querySelector('[name="email"]').value,
			password: form.querySelector('[name="password"]').value,
		}

		try {
			const response = await fetch(url, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(formData),
			})
			console.log(response)

			const data = await response.json()

			if (!response.ok) throw new Error(data.message || 'Ошибка запроса')

			alert(
				url.includes('register')
					? 'Регистрация успешна!'
					: `Добро пожаловать, ${data.user.name}!`
			)

			form.reset()
			if (url.includes('register')) togglePanel(false)
		} catch (error) {
			alert(`Ошибка: ${error.message}`)
		}
	}

	document.getElementById('loginForm').addEventListener('submit', e => {
		e.preventDefault()
		handleForm(e.target, '/api/login')
	})

	document.getElementById('registerForm').addEventListener('submit', e => {
		e.preventDefault()
		handleForm(e.target, '/api/register')
	})

	// Инициализация
	updateMobileSwitcher()
})
