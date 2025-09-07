// Function to generate preview pathname based on content type and document
const getPreviewPathname = (uid, {locale, document}): string => {
    const {slug} = document;
    // Handle different content types with their specific URL patterns
    switch (uid) {
        case "api::project.project":
            console.log(slug)
            break;
    }
    return "";
}

function getPreviewConfig(env) {
    const clientUrl = env("CLIENT_URL"); // Frontend application URL
    const previewSecret = env("PREVIEW_SECRET"); // Secret key for preview authentication

    return {
        enabled: true,
        config: {
            allowedOrigins: clientUrl,
            async handler(uid, {documentId, locale, status}) {
                // Fetch the complete document from Strapi
                const document = await strapi.documents(uid).findOne({documentId});

                // Generate the preview pathname based on content type and document
                const pathname = getPreviewPathname(uid, {locale, document});

                // Disable preview if the pathname is not found
                if (!pathname) {
                    return null;
                }

                // Use Next.js draft mode passing it a secret key and the content-type status
                const urlSearchParams = new URLSearchParams({
                    url: pathname,
                    secret: previewSecret,
                    status,
                });
                return `${clientUrl}?${urlSearchParams}`;
            },
        }
    }
}

export default ({env}) => ({
    auth: {
        secret: env('ADMIN_JWT_SECRET'),
    },
    apiToken: {
        salt: env('API_TOKEN_SALT'),
    },
    transfer: {
        token: {
            salt: env('TRANSFER_TOKEN_SALT'),
        },
    },
    secrets: {
        encryptionKey: env('ENCRYPTION_KEY'),
    },
    flags: {
        nps: env.bool('FLAG_NPS', true),
        promoteEE: env.bool('FLAG_PROMOTE_EE', true),
    },
    preview: getPreviewConfig(env)
})

