const response = await fetch("/static/files.json");
if (!response.ok) {
  throw new Error(`Failed to load JSON: ${response.status}`);
}
const files = await response.json();

const miniSearch = new MiniSearch({
  fields: ["title", "excerpt", "content", "route"], // fields to index for full-text search
  storeFields: ["title", "excerpt", "content", "route"], // fields to return with search results
});

miniSearch.addAll(files);

function search(query) {
  const result = miniSearch.search(query);
  console.log(result);
  return result;
}

const url = new URL(window.location.href);
const searchParams = new URLSearchParams(url.search);
const q = searchParams.get("q");

const result = search(q || "");

const searchTarget = document.getElementById("search-result");

for (const file of result) {
  const a = document.createElement("a");
  a.href = file.route;
  a.textContent = file.title;

  const h2 = document.createElement("h2");
  h2.appendChild(a);

  const li = document.createElement("li");
  li.appendChild(h2);

  const div1 = document.createElement("div");
  div1.textContent = file.excerpt;

  const a2 = document.createElement("a");
  a2.href = file.route;
  a2.textContent = "Read more";

  const div2 = document.createElement("div");
  div2.appendChild(a2);
  div2.classList.add("readmore");

  li.appendChild(div1);
  li.appendChild(div2);

  searchTarget.appendChild(li);
}
