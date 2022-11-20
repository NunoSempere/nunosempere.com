## Description:

## Libraries

### Install
# install.packages("ggplot2")
#install.packages("readr")

### Load
library("ggplot2")
library("readr")
library("ggthemes")
library("magrittr")
library("RColorBrewer")
library("ggsci")


## Data import
setwd("/home/loki/Documents/core/ea/fresh/misc/openphil-funding")
data <- read.csv("grants.csv", header=TRUE, stringsAsFactors = FALSE)

## Data cleaning
colnames(data)
getYear <- function(dateRow){
  year = strsplit(dateRow, " ")[[1]][2]
  return(year)
}
getYear(data$Date[1])
as.vector(sapply(data$Date, getYear))
df <- list()
df$year <- as.vector(sapply(data$Date, getYear))
df$amount <- as.vector(sapply(data$Amount, parse_number))
df$amount <- ifelse(is.na(df$amount), 0, df$amount)
df$area <- as.vector(data$Focus.Area)

df <- as.data.frame(df)
## View(df)
df$area <- as.vector(data$Focus.Area)

areas <- unique(df$area)

ea_growth <-  c("Effective Altruism Community Growth", "Effective Altruism Community Growth (Global Health and Wellbeing)")
global_health <- c("South Asian Air Quality", "Human Health and Wellbeing", "GiveWell-Recommended Charities", "Global Aid Policy", "Global Health & Wellbeing", "Global Health & Development","Science for Global Health")
longtermism <- c("Biosecurity & Pandemic Preparedness", "Potential Risks from Advanced AI", "Science Supporting Biosecurity and Pandemic Preparedness", "Longtermism")
animal_welfare <- c("Farm Animal Welfare", "Broiler Chicken Welfare", "Cage-Free Reforms", "Alternatives to Animal Products")
scientific_research <- c("Transformative Basic Science", "Scientific Research", "Other Scientific Research Areas", "Scientific Innovation: Tools and Techniques")
politicy_advocacy <- c("Land Use Reform","Macroeconomic Stabilization Policy", "Criminal Justice Reform", "Immigration Policy")
not_other <- c(ea_growth, global_health, longtermism, animal_welfare, scientific_research, politicy_advocacy)
other <- areas[!(areas %in% not_other)]

df$area <- ifelse(df$area %in% ea_growth, "EA Community Building", df$area)
df$area <- ifelse(df$area %in% global_health, "Global Health and Wellbeing", df$area)
df$area <- ifelse(df$area %in% longtermism, "Longtermism & GCRs", df$area)
df$area <- ifelse(df$area %in% animal_welfare, "Animal Welfare", df$area)
df$area <- ifelse(df$area %in% scientific_research, "Scientific Research", df$area)
df$area <- ifelse(df$area %in% politicy_advocacy, "Policy Advocacy", df$area)
df$area <- ifelse(df$area %in% other, "Other", df$area)
df$area
# df <- df[order(df$area),]

## Aggregate by year and area
years <- c(2014: 2022)# as.vector(unique(df$year))
num_years <- length(years)
area_names <- as.vector(unique(df$area))
num_areas <- length(area_names)

df2 <- list()
df2$area <- sort(rep(area_names, num_years))
df2$year <- rep(years, num_areas)
df2 <- as.data.frame(df2)
 
# View(df2)

getAmountForYearAreaPair <- function(target_year, target_area){
  filter = dplyr::filter
  # target_year = 2022
  # target_area = "Longtermism"
  rows = df %>% filter(year == target_year) %>% filter(area == target_area)
  return(sum(rows$amount))
}
getAmountForYearAreaPair(2022, "Longtermism")

amounts <- c()
for(i in c(1:dim(df2)[1])){
  amount <- getAmountForYearAreaPair(df2$year[i], df2$area[i])
  amounts <- c(amounts, amount)
}
df2$amount <- amounts
# View(df2)

## Plotting
title_text="Open Philanthropy allocation by year and cause area"
subtitle_text="with my own aggregation of categories"
palette = "Classic Red-Blue"
direction = -1
open_philanthropy_plot <- ggplot(data=df2, aes(x=year, y=amount, fill=area, group=amount))+
  geom_bar(stat="identity")+
  labs(
    title=title_text,
    subtitle=subtitle_text,
    x=element_blank(),
    y=element_blank()
  ) +
  # scale_fill_wsj() +
  # scale_fill_tableau(dir =1) +
  # scale_fill_tableau(palette, dir=direction) +
  # scale_fill_viridis(discrete = TRUE) +
  # scale_fill_brewer(palette = "Set2") +
  
  scale_fill_d3( "category20", alpha=0.8) +
  # scale_fill_uchicago("dark") +
  # scale_fill_startrek() +
  scale_y_continuous(labels = scales::dollar_format(scale = 0.000001, suffix = "M"), breaks = c(0:6)*10^8)+
  scale_x_continuous(breaks = years)+
  theme_tufte() +
  theme(
    legend.title = element_blank(),
    plot.title = element_text(hjust = 0.5),
    plot.subtitle = element_text(hjust = 0.5),
    legend.position="bottom",
    legend.box="vertical",
    axis.text.x=element_text(angle=60, hjust=1),
    legend.text=element_text(size=7, hjust = 0.5)
  ) + 
  guides(fill=guide_legend(nrow=3,byrow=TRUE))
getwd() ## Working directory on which the file will be saved. Can be changed with setwd("/your/directory")
height = 5
width = 5
ggsave(plot=open_philanthropy_plot, "open_philanthropy_grants_stacked.png", width=width, height=height, bg = "white")

## Including Dustin Moskovitz's wealth 
coeff <- 10^7*4
wealth <- c(6, 8, 12, 15, 18, 12, 14, 19, 14)
df2$wealth <- rep(wealth * coeff, num_areas)

open_philanthropy_plot_with_fortune <- ggplot(data=df2, aes(x=year, y=amount, fill=area, group=amount))+
  geom_bar(stat="identity")+
  geom_point(aes(x=year, y=wealth), size=2, color="darkblue", shape=4)+
  labs(
    title=title_text,
    subtitle=subtitle_text,
    x=element_blank(),
    y=element_blank()
  ) +
  # scale_fill_wsj() +
  # scale_fill_tableau(dir =1) +
  # scale_fill_tableau(palette, dir=direction) +
  # scale_fill_viridis(discrete = TRUE) +
  # scale_fill_brewer(palette = "Set2") +
  
  scale_fill_d3( "category20", alpha=0.8) +
  # scale_fill_uchicago("dark") +
  # scale_fill_startrek() +
  scale_y_continuous(
    labels = scales::dollar_format(scale = 0.000001, suffix = "M"), 
    name="OpenPhil donations", 
    breaks = c(0:5)*10^8, 
    sec.axis = sec_axis(
      ~.*1, 
      name="Dustin Moskovitz's fortune\n(est. Bloomberg)", 
      breaks = seq(0,20,by=5)*coeff,
      labels = c("$0B", "$5B","$10B","$15B", "$20B")
      ),
    limits=c(0,8*10^8)
    )+
  scale_x_continuous(breaks = years)+
  theme_tufte() +
  theme(
    legend.title = element_blank(),
    plot.title = element_text(hjust = 0.5),
    plot.subtitle = element_text(hjust = 0.5),
    legend.position="bottom",
    legend.box="vertical",
    axis.text.x=element_text(angle=60, hjust=1),
    axis.title.y = element_text(vjust=3, hjust=0.25, size=10),
    axis.title.y.right = element_text(vjust=3, hjust=0.5, size=10),
    legend.text=element_text(size=8)
  ) + 
  guides(fill=guide_legend(nrow=4,byrow=TRUE))
# open_philanthropy_plot_with_fortune

height = 6
width = 5

ggsave(plot=open_philanthropy_plot_with_fortune, "open_philanthropy_plot_with_fortune.png", width=width, height=height, bg = "white")