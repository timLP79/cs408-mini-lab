# Canvas Module Progress CLI

A command-line tool written in Go that connects to the Canvas LMS REST API and displays module completion progress for your courses. Includes full-text search across module item titles and Canvas Page content. Useful for quickly checking which modules you've completed or finding where specific topics are covered without opening a browser.

## Demo

![Demo](assets/demo.gif)

## Setup Instructions

### Prerequisites

- [Go 1.21+](https://go.dev/dl/)
- A Canvas API token (see steps below)

### 1. Clone the repo

```bash
git clone https://github.com/timLP79/cs408-mini-lab.git
cd cs408-mini-lab
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Create your `.env` file

Copy the example file and fill in your credentials:

```bash
cp .env.example .env
```

Open `.env` and replace the placeholder values:

```
CANVAS_API_TOKEN=your_token_here
CANVAS_BASE_URL=https://boisestatecanvas.instructure.com
```

**How to get your Canvas API token:**
1. Log in to Canvas and click your profile picture in the left sidebar
2. Click **Settings**
3. Scroll to **Approved Integrations** and click **+ New Access Token**
4. Set a purpose (e.g. "CS408 CLI") and an expiry date, then click **Generate Token**
5. Copy the token immediately. Canvas will only show it once

### 4. Run the tool

```bash
go run .
```

## Example Usage

### List your courses

```bash
go run .
```

Expected output:
```
Your Courses:
-------------------------------
1: COEN Undergrad Students
2: Department of Computer Science - Students Groups
3: Fa24 - CS 153 - Navigating Computer Systems
4: Sp26 - CS 408 - Full Stack Web Development
5: Sp26 - CS 410/510 - Databases
6: Sp26 - MATH 301 - Introduction to Linear Algebra

Enter course number:
```

### View module progress for a course

```
Enter course number: 6

Sp26 - CS 408 - Full Stack Web Development
------------------------------
[ ] Course Resources
[~] Week 1 - Introduction and Overview   [████████--] 8/10
[~] Week 2 - CS208 Database Review       [██--] 2/4
[✓] Week 3 - Tech Stack                  [██] 2/2
[✓] Week 4 - Form Teams                  [███] 3/3
[ ] Week 5 - Developer Setup
[✓] Week 6 - AWS                         [███] 3/3
[✓] Week 7 - Project Specification       [█] 1/1
[ ] Week 8
[ ] Week 16 - Final Project Showcase     [---] 0/3
```

Column width adjusts dynamically to the longest module name in the selected course. Progress bars scale to the number of trackable items in each module, each character represents one item.

### Search module content

After viewing module progress, enter a search query to find items by title or page content:

```
Enter a search query: matrix multiplication

Search results for "matrix multiplication"
-------------------------------------------
[Week 7: Matrix Arithmetic and Inverses]   [Week 7 Overview]   (body)
	...ntroduction
This week, we will study operations with matrices (addition, scalar/matrix multiplication, and inverse)...
```

Search checks item titles first. If a title matches, the page body is skipped to avoid duplicate results. For Page items with no title match, the full page body is searched and a snippet of surrounding context is shown.

**Module status legend:**
- `[✓]` green - completed (with a tracked completion timestamp)
- `[~]` yellow - in progress
- `[ ]` white - unlocked but not started, or no trackable items
- `[🔒]` red - locked

## API Endpoints Used

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/v1/courses` | GET | Retrieves all courses the authenticated user is enrolled in |
| `/api/v1/courses/:id/modules` | GET | Retrieves all modules and their completion state for a given course |
| `/api/v1/courses/:id/modules/:id/items` | GET | Retrieves individual items within a module and their completion status |
| `/api/v1/courses/:id/pages/:page_url` | GET | Retrieves the full content body of a Canvas Page item for search |

All endpoints handle pagination by following the `Link` response header until all pages are retrieved.

## Reflection
This was a very interesting project for me. In my current role at work, I am an implementation engineer. So it is my job to help companies connect to our API for the services that 
we provide. This project combined with the semester long project really help give me a better understanding of what other companies are going through and then being able to actually 
build my own api routes has been a great lesson as well. I used curl several times throughout this process to see what fields were available in the API endpoints and I think I can apply this
to my projects at work to help customers in a more technical way. This definitely was a fun project and taught me a lot.
  
The challenging part of the project was getting the results to display exactly as I wanted them to. Some items in the modules 
don't actually have items which can be completed and so my counts and results were not quite correct. I had to go back and add a separate counter which 
counted the items and then one which actually counted items which were actually trackable.

If I had more time, I would really polish the search feature. I added this functionality at the last minute. It works to an extent, but is not as robust as I would like. That is the main reason 
that I did not include this in my GIF. The main functionality is querying the modules and giving a display of progress at a module level. Also, I would also track the dates of the items to make 
that part of the logic so that items might be color coded based on if they are current, past, or future items.