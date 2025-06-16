# preview


![image](https://github.com/user-attachments/assets/2389310e-234d-441c-ab2e-273fce897157)


--- 
# how to add a new entry to the website? 

1. make a change to the data/members.toml file.
2. add your entry/details
3. commit the change

--- 
# how do i use the webring in my portfolio? 

```html
<a href="https://iitm-paradox.github.io/webring/YOUR_SLUG/previous">&larr;</a>
<a href="https://iitm-paradox.github.io/webring/">WebOps Webring</a>
<a href="https://iitm-paradox.github.io/webring/YOUR_SLUG/next">&rarr;</a>
```


<details>
<summary>styling the webring?</summary>

```html
<!-- WebOps Webring Badge -->
<div style="text-align: center; margin: 2rem 0; font-size: 0.95rem; color: gray;">
  <a href="https://iitm-paradox.github.io/webring/YOUR_SLUG/previous" target="_blank" style="margin: 0 0.5rem; text-decoration: none;">&larr;</a>
  <a href="https://iitm-paradox.github.io/webring/" target="_blank" style="margin: 0 0.5rem; text-decoration: none;">WebOps Webring</a>
  <a href="https://iitm-paradox.github.io/webring/YOUR_SLUG/next" target="_blank" style="margin: 0 0.5rem; text-decoration: none;">&rarr;</a>
</div>
   
