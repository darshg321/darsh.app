import { Marked } from 'marked';

const plain = new Marked({ gfm: true, breaks: false });

/** short markdown from the data files */
export const md = (src: string) => plain.parse(src, { async: false }) as string;

/** a build journal: relative images live in the GitHub repo, and everything loads lazily */
export function journal(src: string, repo: string) {
	const raw = `https://raw.githubusercontent.com/${repo}/HEAD/`;
	const m = new Marked({
		gfm: true,
		breaks: true,
		walkTokens(token) {
			if ((token.type === 'image' || token.type === 'link') && !/^(https?:|mailto:|#|\/)/.test(token.href)) {
				token.href = raw + token.href.replace(/^\.\//, '');
			}
		},
		renderer: {
			image({ href, text }) {
				return `<img src="${href}" alt="${text.replace(/"/g, '&quot;')}" loading="lazy" decoding="async">`;
			},
			link({ href, tokens }) {
				return `<a href="${href}" target="_blank" rel="noreferrer">${this.parser.parseInline(tokens)}</a>`;
			}
		}
	});
	// the Blueprint export starts with an HTML comment and some journals have raw <img> tags
	return m.parse(src.replace(/<!--[\s\S]*?-->/g, ''), { async: false }) as string;
}
