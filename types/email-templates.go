package types

type EmailTemplate string

const (
	EmailTemplateForgotPassword   EmailTemplate = "forgot"
	EmailTemplateValidateAccount  EmailTemplate = "validate-account"
	EmailTemplateRegisterSuccess  EmailTemplate = "register_success"
	EmailTemplateNewInvoiceCall   EmailTemplate = "new-invoice-call"
	EmailTemplateExportComparator EmailTemplate = "export-comparator"
	EmailTemplateFeedSynchro      EmailTemplate = "feed-synchro"
	EmailTemplateFeedGeneration   EmailTemplate = "feed-generation"
)
