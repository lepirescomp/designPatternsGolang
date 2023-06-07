package main

import "fmt"

type email struct {
	from, to, subject string
}

type EmailBuilder struct {
	e email
}

func (b *EmailBuilder) BuildFrom(s string) *EmailBuilder {
	b.e.from = s
	return b
}

func (b *EmailBuilder) BuildTo(s string) *EmailBuilder {
	b.e.to = s
	return b
}

func (b *EmailBuilder) BuildSubject(s string) *EmailBuilder {
	b.e.subject = s
	return b
}

func sendEmail(e email) {
	fmt.Println("Sending email to", e.to)
}

type build func(*EmailBuilder)

func sendEmailWrapper(action build) {
	b := EmailBuilder{}
	action(&b)
	sendEmail(b.e)
}

func builderEncapsulated() {
	fmt.Println("Builder with control of access")
	sendEmailWrapper(func(builder *EmailBuilder) {
		builder.BuildFrom("from").BuildTo("to").BuildSubject("subject")
	})

}
