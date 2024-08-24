/**
 * Client side search with MiniSearch
 */

(async function main() {
  const files = await getFiles();
  const miniSearch = getMiniSearch(files);
  const query = getQuery();
  if (!query) {
    return;
  }

  const results = search(query, miniSearch);

  const input = getInput();
  input.value = query;

  const routes = getRoutes(results);
  const articles = getArticles();
  filterArticles(articles, routes);
})();

async function getFiles() {
  const response = await fetch("/static/files.json");
  if (!response.ok) {
    throw new Error(`Failed to load JSON: ${response.status}`);
  }
  return await response.json();
}
const files = await getFiles();

function getMiniSearch(files) {
  const miniSearch = new MiniSearch({
    fields: ["title", "excerpt", "content", "route"], // fields to index for full-text search
    storeFields: ["title", "excerpt", "content", "route"], // fields to return with search results
  });

  miniSearch.addAll(files);
  return miniSearch;
}

function search(query, miniSearch) {
  const result = miniSearch.search(query);
  return result;
}

function getQuery() {
  const url = new URL(window.location.href);
  const searchParams = new URLSearchParams(url.search);
  const query = searchParams.get("search");
  return query;
}

function getInput() {
  const form = document.querySelector("#search-form");
  const input = form.querySelector("input[name=search]");
  return input;
}
// input.value = getQuery(); // ?

function getRoutes(results) {
  const routes = results.map((result) => result.route);
  return routes;
}

function getArticles() {
  const articles = document.querySelectorAll(".article-list li");
  const arr = Array.from(articles);
  return arr;
}

function filterArticles(articles, routes) {
  articles
    .filter((atricle) => {
      const href = atricle.querySelector("a")?.href || "";
      const url = new URL(href);
      const pathname = url.pathname;
      if (routes.includes(pathname)) {
        return false;
      } else {
        return true;
      }
    })
    .forEach((atricle) => {
      atricle.style.display = "none";
    });
}
