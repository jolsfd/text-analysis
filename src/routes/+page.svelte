<script>
  import { Header, Content, Grid, Row, Column, Form, TextArea, Button, FormItem  } from "carbon-components-svelte";
  import "./styles.css"

  let text = "";
  let computed = false;

  let res = null;
  let numSentences = 0;
  let numWords = 0;
  let numSyllables = 0;
  let score = 0;
  let readingLevel = "";

  function getReadingLevel(score){
    if (score >= 90){
      return "5th grade"
    }

    if (score >= 80){
      return "6th grade"
    }

    if (score >= 70){
      return "7th grade"
    }

    if (score >= 60){
      return "8-9th grade"
    }

    if (score >= 50){
      return "10-12th grade"
    }

    if (score >= 30){
      return"College Student"
    }

    if (score >= 10){
      return "College Graduate"
    }

    if (score >= 0){
      return "Professional"
    }
  }

  function submit(e){
    e.preventDefault()
    computed = false

    if (text == ""){
      return
    }

    res = computeReadability(text)

    score = Math.round(res[0])
    numWords = res[1]
    numSentences = res[2]
    numSyllables = res[3]

    console.log(score)

    if (!score){
      return
    }

    readingLevel = getReadingLevel(score)
    computed = true
  }
</script>

<Header company="Jolsfd" platformName="Text Analysis" href="/"></Header>

<Content>
  <Grid>
    <Row>
      <Column>
        <h1>Flesch–Kincaid Readability Test</h1>
      </Column>
    </Row>

    <!-- Data input -->
    <Row>
      <Column>
        <Form
          on:submit={(e) => {
            e.preventDefault();
            //console.log("submit", e);
          }}
        >
          <TextArea
            cols={50}
            helperText="Texts are not sent to a cloud provider. Instead, they are analyzed locally on your device."
            invalidText="Invalid error message."
            labelText="Input your text here."
            placeholder="Text to be analyzed..."
            rows={20}
            bind:value={text}
          />

          <div style="padding: var(--cds-spacing-01)"></div>

          <Button size="md" type="submit" on:click={submit}>Analyze</Button>
        </Form>
      </Column>
    </Row>
  </Grid>

  <div style="padding: var(--cds-spacing-05)"></div>

  <!-- Result -->
  {#if computed}
  <Grid>
    <Row>
      <Column>
        <h3>Results</h3>
      </Column>
    </Row>

    <div style="padding: var(--cds-spacing-03)"></div>

    <Row>
      <Column lg={4} md={2} sm={1}>
        <span class="heading-02">Score</span>
      </Column>
      <Column lg={4} md={2} sm={1}>
        <span class="heading-02">Reading Level</span>
      </Column>
    </Row>

    <div style="padding: var(--cds-spacing-02)"></div>

    <Row>
      <Column lg={4} md={2} sm={1}>
        <span>{score}</span>
      </Column>
      <Column lg={4} md={2} sm={1}>
        <span>{readingLevel}</span>
      </Column>
    </Row>

    <div style="padding: var(--cds-spacing-03)"></div>

    <Row>
      <Column lg={4} md={2} sm={1}>
        <span class="text-secondary">Sentences</span>
      </Column>
      <Column lg={4} md={2} sm={1}>
        <span class="text-secondary">Words</span>
      </Column>
      <Column lg={4} md={2} sm={1}>
        <span class="text-secondary">Syllabes</span>
      </Column>
    </Row>
  </Grid>

  <div style="padding: var(--cds-spacing-02)"></div>

  <Grid>
      <Row>
        <Column lg={4} md={2} sm={1}>
          <span>{numSentences}</span>
        </Column>
        <Column lg={4} md={2} sm={1}>
          <span>{numWords}</span>
        </Column>
        <Column lg={4} md={2} sm={1}>
          <span>{numSyllables}</span>
        </Column>
      </Row>
  </Grid>
  {/if}
</Content>

<style>
  .text-secondary {
    color: var(--cds-text-secondary);
  }

  .heading-02{
    letter-spacing: var(--cds-productive-heading-02-letter-spacing);
    line-height: var(--cds-productive-heading-02-line-height);
    font-weight: var(--cds-productive-heading-02-font-weight);
    font-size: var(--cds-productive-heading-02-font-size);
  }
</style>