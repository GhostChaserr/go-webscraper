

-- Link table
CREATE TABLE IF NOT EXISTS webscraper.links_raw (
  `link_id`      UUID DEFAULT generateUUIDv4(),
  `document_id`  String,
  `link`         String DEFAULT ''
)

ENGINE = MergeTree()
ORDER BY link_id;

CREATE TABLE IF NOT EXISTS webscraper.texts_raw (
  `text_id`      UUID DEFAULT generateUUIDv4(),
  `document_id`  String,
  `content`      String DEFAULT ''
)

ENGINE = MergeTree()
ORDER BY text_id;

CREATE TABLE IF NOT EXISTS webscraper.og_tags_raw (
  `tag_id`        UUID DEFAULT generateUUIDv4(),
  `document_id`   String,
  `type`          String DEFAULT '',
  `title`         String DEFAULT '',
  `description`   String DEFAULT '',
  `site_name`     String DEFAULT '',
  `url`           String DEFAULT '',
  `image`         String DEFAULT ''
)

ENGINE = MergeTree()
ORDER BY tag_id;

CREATE TABLE IF NOT EXISTS webscraper.twitter_tags_raw (
  `tag_id`        UUID DEFAULT generateUUIDv4(),
  `document_id`   String,
  `title`         String DEFAULT '',
  `description`   String DEFAULT '',
  `card`          String DEFAULT '',
  `domain`        String DEFAULT '',
  `image`         String DEFAULT ''
)

ENGINE = MergeTree()
ORDER BY tag_id;

CREATE TABLE IF NOT EXISTS webscraper.keywords_raw (
  `keyword_id`        UUID DEFAULT generateUUIDv4(),
  `document_id`       String,
  `keyword`           String DEFAULT ''
)

ENGINE = MergeTree()
ORDER BY keyword_id;

CREATE TABLE IF NOT EXISTS webscraper.images_raw (
  `image_id`        UUID DEFAULT generateUUIDv4(),
  `document_id`     String,
  `image`           String DEFAULT ''
)

ENGINE = MergeTree()
ORDER BY image_id;


-- Main document table
CREATE TABLE IF NOT EXISTS webscraper.document_raw (
  `created_at`                 DateTime DEFAULT now(),
  `document_id`                UUID DEFAULT generateUUIDv4(),
  `heading`                    String DEFAULT '',
  `language`                   String DEFAULT '',
  `words_count`                UInt16 DEFAULT 0,
  `title`                      String DEFAULT '',
  `description`                String DEFAULT '',
  `url`                        String DEFAULT '',
  `canonical`                  String DEFAULT '',
  `fav_icon`                   String DEFAULT '',
  `manifest`                   String DEFAULT '',
  `internal_scripts_count`     UInt16 DEFAULT 0,
  `external_scripts_count`     UInt16 DEFAULT 0,
  `outbound_links_count`       UInt16 DEFAULT 0,
  `inbound_links_count`        UInt16 DEFAULT 0,
  `external_css_count`         UInt16 DEFAULT 0,
  `internal_css_count`         UInt16 DEFAULT 0,
  `has_viewport_tag`           Boolean DEFAULT 1,
  `has_gsiteverification_tag`  Boolean DEFAULT 1
)

ENGINE = MergeTree()
PARTITION BY toYYYYMM(created_at)
ORDER BY document_id;