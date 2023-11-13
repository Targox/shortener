import type { PageServerLoad, Actions } from './$types'
import { fail } from '@sveltejs/kit'
import { superValidate } from 'sveltekit-superforms/server'
import { formSchema } from './schema'
import { db } from '$lib/db'
import { user as userSchema } from '$lib/db/schema'
import { eq } from 'drizzle-orm'
import { nanoid } from 'nanoid'

export const load = (async () => {
	return {
		form: superValidate(formSchema),
	}
}) satisfies PageServerLoad

export const actions: Actions = {
	default: async (event) => {
		const form = await superValidate(event, formSchema)

		if (!form.valid) {
			return fail(400, {
				form,
			})
		}

		try {
			const users = await db
				.select()
				.from(userSchema)
				.where(eq(userSchema.email, form.data.email))

			const user = users[0]
			const matchPassword =
				user && (await Bun.password.verify(form.data.password, user.password))

			if (user && matchPassword) {
				const token = nanoid(32)
				event.cookies.set('token', token, {
					httpOnly: true,
					sameSite: 'strict',
					path: '/',
				})
				return {
					form,
				}
			} else {
				return fail(400, {
					form,
				})
			}
		} catch (error) {
			if (error instanceof SyntaxError) {
				return fail(400, {
					form,
				})
			} else {
				console.log(error)
				return fail(400, {
					form,
				})
			}
		}
	},
}
